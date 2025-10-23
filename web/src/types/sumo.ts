export type Language = 'en' | 'jp'

export const MATCH_RESULT = {
  PENDING: -1,
  LOST: 0,
  WON: 1,
} as const

export type MatchResult = typeof MATCH_RESULT[keyof typeof MATCH_RESULT]

export const MATCH_SIDE = {
  EAST: 'east',
  WEST: 'west',
} as const

export type MatchSide = typeof MATCH_SIDE[keyof typeof MATCH_SIDE]

export interface RikishiProfileDetails {
  stable: string
  stable_url: string
  name: string
  ring_name: string
  current_rank: string
  birthday: string
  birthplace: string
  height: string
  weight: string
  signature_maneuver: string
  image_url: string
}

export interface RikishiProfile {
  id: string
  profiles: {
    en: RikishiProfileDetails
    jp: RikishiProfileDetails
  }
}

export interface Rikishi {
  id: string
  shikona_jp: string
  shikona_en: string
  banzuke_name_jp: string
  banzuke_name_en: string
  won: number
  lost: number
  rest_jp: string
  rest_en: string
  result: MatchResult
  kyokai_member_id: string
  profile?: RikishiProfile
}

export interface Match {
  east: Rikishi
  west: Rikishi
  judge_id: string
  technic_id: string
}
