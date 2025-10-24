import type { ConnectionState } from '@/services/websocket'

interface WebSocketStatusIndicatorProps {
  isConnected: boolean
  connectionState: ConnectionState
  error: Error | null
  channelKey?: string
  showDebugInfo?: boolean
}

export default function WebSocketStatusIndicator({
  isConnected,
  connectionState,
  error,
  channelKey,
  showDebugInfo = false
}: WebSocketStatusIndicatorProps) {
  if (isConnected) {
    return (
      <span
        className="ml-2 inline-block w-2 h-2 bg-green-500 rounded-full"
        title={channelKey ? `Connected to ${channelKey}` : 'Connected'}
      />
    )
  }

  if (connectionState === 'connecting') {
    return (
      <span
        className="ml-2 inline-block w-2 h-2 bg-yellow-500 rounded-full animate-pulse"
        title="Connecting..."
      />
    )
  }

  if (error || connectionState === 'error') {
    return (
      <span
        className="ml-2 inline-block w-2 h-2 bg-red-500 rounded-full"
        title={error ? `Error: ${error.message}` : 'Connection error'}
      />
    )
  }

  if (connectionState === 'disconnected') {
    return showDebugInfo ? (
      <span
        className="ml-2 inline-block w-2 h-2 bg-gray-400 rounded-full"
        title="Disconnected"
      />
    ) : null
  }

  return null
}
