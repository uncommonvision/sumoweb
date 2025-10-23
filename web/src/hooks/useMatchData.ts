import { useState, useEffect } from 'react'
import { matchService } from '@/services/matches'
import type { Match } from '@/types'

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

  return { matches, loading, error }
}
