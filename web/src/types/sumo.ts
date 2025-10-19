export type Language = 'en' | 'jp'

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
  result: boolean
  kyokai_member_id: string
}

export interface Match {
  east: Rikishi
  west: Rikishi
  judge_id: string
  technic_id: string
}
