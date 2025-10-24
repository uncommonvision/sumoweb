import { useEffect, useState, useCallback, useRef } from 'react'
import { websocketService, handleConnectionAck, type ConnectionState } from '@/services/websocket'
import type { WebsocketSumoMessage, MessagePayload } from '@/types'

interface UseMatchSocketServiceOptions {
  division: string
  day: string
  autoConnect?: boolean
  onConnectionChange?: (connected: boolean, state: ConnectionState) => void
}

interface UseMatchSocketServiceReturn {
  isConnected: boolean
  connectionState: ConnectionState
  error: Error | null
  channelKey: string
  sendMessage: (message: Omit<MessagePayload, 'id' | 'sendAt'>) => void
  onMessage: (callback: (message: WebsocketSumoMessage) => void) => () => void
  reconnect: () => Promise<void>
  disconnect: () => void
}

export function useMatchSocketService(options: UseMatchSocketServiceOptions): UseMatchSocketServiceReturn {
  const { division, day, autoConnect = true, onConnectionChange } = options

  const [error, setError] = useState<Error | null>(null)
  const [, forceUpdate] = useState({})

  const onConnectionChangeRef = useRef(onConnectionChange)

  const channelKey = `sumo:division:${division}:day:${day}`

  useEffect(() => {
    onConnectionChangeRef.current = onConnectionChange
  }, [onConnectionChange])

  useEffect(() => {
    const unsubscribe = websocketService.onStateChange((state) => {
      onConnectionChangeRef.current?.(state === 'connected', state)
      forceUpdate({})
    })

    if (autoConnect && division && day) {
      const connect = async () => {
        try {
          setError(null)
          await websocketService.connect(channelKey)
        } catch (err) {
          const error = err instanceof Error ? err : new Error('Failed to connect')
          setError(error)
        }
      }
      connect()
    }

    return unsubscribe
  }, [channelKey, division, day, autoConnect])

  const connectionState = websocketService.getConnectionState()
  const isConnected = websocketService.isConnected()

  useEffect(() => {
    const unsubscribeAck = handleConnectionAck(() => {
      console.log('Match socket connected:', channelKey)
    })

    return unsubscribeAck
  }, [channelKey])

  const reconnect = useCallback(async () => {
    try {
      setError(null)
      await websocketService.connect(channelKey)
    } catch (err) {
      const error = err instanceof Error ? err : new Error('Failed to reconnect')
      setError(error)
    }
  }, [channelKey])

  const disconnect = useCallback(() => {
    websocketService.disconnect()
  }, [])

  const sendMessage = useCallback((message: Omit<MessagePayload, 'id' | 'sendAt'>) => {
    if (!websocketService.isConnected()) {
      console.warn('Cannot send message: WebSocket not connected')
      return
    }
    
    const fullMessage: MessagePayload = {
      ...message,
      id: crypto.randomUUID(),
      sentAt: new Date().toISOString()
    }
    
    websocketService.send('CHAT_MESSAGE', fullMessage)
  }, [])

  const onMessage = useCallback((callback: (message: WebsocketSumoMessage) => void) => {
    const unsubscribe = websocketService.on('MATCH_UPDATE', callback)
    return unsubscribe
  }, [])

  return {
    isConnected,
    connectionState,
    error,
    channelKey,
    sendMessage,
    onMessage,
    reconnect,
    disconnect
  }
}
