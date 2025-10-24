import { useState, useEffect } from 'react'
import { matchService } from '@/services/matches'
import { websocketService } from '@/services/websocket'
import type { Match } from '@/types'
import type { MatchUpdatePayload } from '@/types/websocket'

export function useMatchData(division?: string, day?: string) {
  const [matches, setMatches] = useState<Match[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (!division || !day) {
      setLoading(false)
      return
    }

    const divisionNum = parseInt(division, 10)
    const dayNum = parseInt(day, 10)

    if (isNaN(divisionNum) || isNaN(dayNum)) {
      setError('Invalid division or day parameter')
      setLoading(false)
      return
    }

    let cancelled = false

    async function fetchData() {
      try {
        setLoading(true)
        setError(null)

        const data = await matchService.fetchMatches(divisionNum, dayNum)

        if (!cancelled) {
          setMatches(data)
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : 'Failed to fetch matches')
        }
      } finally {
        if (!cancelled) {
          setLoading(false)
        }
      }
    }

    fetchData()

    return () => {
      cancelled = true
    }
  }, [division, day])

  useEffect(() => {
    if (!division || !day) return

    const divisionNum = parseInt(division, 10)
    const dayNum = parseInt(day, 10)

    if (isNaN(divisionNum) || isNaN(dayNum)) return

    const unsubscribe = websocketService.on<MatchUpdatePayload>('MATCH_UPDATE', (payload) => {
      if (payload.division === divisionNum && payload.day === dayNum) {
        console.log('Received MATCH_UPDATE:', payload)
        
        setMatches((prevMatches) => {
          const mergedMatches = matchService.mergeMatches(prevMatches, payload.matches)
          const updatedCount = payload.matches?.length || 0
          const totalCount = prevMatches.length
          
          console.log(`Match merge: ${updatedCount} updates received, ${totalCount} total matches, ${mergedMatches.length} after merge`)
          
          return mergedMatches
        })
      }
    })

    return unsubscribe
  }, [division, day])

  return { matches, loading, error }
}
