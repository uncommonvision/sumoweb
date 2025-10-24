import { useParams } from 'react-router-dom'
import { useEffect, useRef } from 'react'
import { NoAuthLayout } from "@/components/layout"
import { MatchList } from "@/containers"
import { WebSocketStatusIndicator } from "@/components/ui"
import { useLanguagePreference } from '@/hooks/useLanguagePreference'
import { useKeydownShortcut } from '@/hooks/useKeydownShortcut'
import { useMatchData } from '@/hooks/useMatchData'
import { useMatchSocketService } from '@/hooks/useMatchSocketService'
import { useUserIdentity } from '@/hooks/useUserIdentity'
import { websocketService } from '@/services/websocket'

export default function SumoPage() {
  const { division, day } = useParams<{ division: string; day: string }>()
  const [language, setLanguage] = useLanguagePreference()
  const { matches, loading, error } = useMatchData(division, day)
  const { getOrCreateUser } = useUserIdentity()
  const hasIdentifiedRef = useRef(false)

  const {
    isConnected: wsConnected,
    connectionState,
    error: wsError,
    channelKey
  } = useMatchSocketService({
    division: division || '',
    day: day || '',
    autoConnect: true,
    onConnectionChange: (connected, state) => {
      console.log('WebSocket connection changed:', { connected, state, channelKey })
    }
  })

  useEffect(() => {
    if (wsConnected && !hasIdentifiedRef.current) {
      const user = getOrCreateUser()
      
      websocketService.send('USER_IDENTIFY', {
        userId: user.id,
        userName: user.name
      })
      
      hasIdentifiedRef.current = true
      console.log('Sent USER_IDENTIFY:', { userId: user.id, userName: user.name })
    }

    if (!wsConnected) {
      hasIdentifiedRef.current = false
    }
  }, [wsConnected, getOrCreateUser])

  useKeydownShortcut(
    { key: 'l', ctrl: false },
    () => setLanguage(language === 'en' ? 'jp' : 'en'),
    'Toggle Language',
    'Switch between English and Japanese display'
  )

  return (
    <NoAuthLayout className="">
      <div className="space-y-6">
        <div className="text-center">
          <h1 className="text-3xl font-bold text-foreground mb-2">
            Sumo Torikumi
            <WebSocketStatusIndicator
              isConnected={wsConnected}
              connectionState={connectionState}
              error={wsError}
              channelKey={channelKey}
              showDebugInfo={import.meta.env.DEV}
            />
          </h1>
          <p className="text-lg text-muted-foreground">
            Division: {division} | Day: {day}
          </p>
          {process.env.NODE_ENV === 'development' && (
            <p className="text-xs text-muted-foreground mt-1">
              WS: {channelKey} ({connectionState})
            </p>
          )}
        </div>

        {loading && (
          <div className="text-center text-muted-foreground">
            Loading matches...
          </div>
        )}

        {error && (
          <div className="text-center text-red-500">
            Error: {error}
          </div>
        )}

        {!loading && !error && <MatchList matches={matches} />}
      </div>
    </NoAuthLayout>
  )
}
