import type { Match } from '@/types'

export const matchService = {
  async fetchMatches(division: number, day: number): Promise<Match[]> {
    const response = await fetch(`/api/v1/matches/${division}/${day}`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch matches: ${response.statusText}`)
    }
    
    return response.json()
  },

  areSameMatch(match1: Match, match2: Match): boolean {
    return match1.east.id === match2.east.id && match1.west.id === match2.west.id
  },

  mergeMatches(currentMatches: Match[], payloadMatches: Match[]): Match[] {
    if (!payloadMatches || payloadMatches.length === 0) {
      return currentMatches
    }

    return currentMatches.map(currentMatch => {
      const updatedMatch = payloadMatches.find(payloadMatch => 
        this.areSameMatch(currentMatch, payloadMatch)
      )
      return updatedMatch || currentMatch
    })
  }
}
