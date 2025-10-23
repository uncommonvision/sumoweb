import type { Match } from '@/types'

export const matchService = {
  async fetchMatches(division: number, day: number): Promise<Match[]> {
    const response = await fetch(`/api/v1/matches/${division}/${day}`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch matches: ${response.statusText}`)
    }
    
    return response.json()
  }
}
