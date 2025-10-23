import type { Rikishi, MatchSide } from '@/types'
import { MATCH_RESULT, MATCH_SIDE } from '@/types'

interface RikishiMatchCardResultProps {
  rikishi: Rikishi
  side: MatchSide
}

export default function RikishiMatchCardResult({ rikishi, side }: RikishiMatchCardResultProps) {
  if (rikishi.result === MATCH_RESULT.PENDING) {
    return null
  }

  return (
    <div className={`absolute bottom-0 bg-stone-500 p-1
      ${side === MATCH_SIDE.EAST ? 'left-0 rounded-tr rounded-bl' : 'right-0 rounded-tl rounded-br'}`}>
      <div className={`w-4 h-4 rounded-full shadow ${rikishi.result === MATCH_RESULT.WON ? 'bg-white' : 'bg-black'}`}></div>
    </div>
  )
}
