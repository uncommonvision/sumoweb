import type { Rikishi } from '@/types'
import RikishiMatchCardImage from '../RikishiMatchCardImage'
import RikishiMatchCardResult from '../RikishiMatchCardResult'
import { useLanguagePreference } from '@/hooks/useLanguagePreference'

interface RikishiMatchCardProps {
  rikishi: Rikishi
  side: 'east' | 'west'
  isExpanded?: boolean
}

export default function RikishiMatchCard({ rikishi, side, isExpanded = false }: RikishiMatchCardProps) {
  const [language] = useLanguagePreference()
  const isEast = side === 'east'
  return (
    <div className="bg-card relative text-card-foreground flex items-start h-full">
      <RikishiMatchCardResult rikishi={rikishi} side={side} />

      <div className={`flex ${isEast ? 'flex-row' : 'flex-row-reverse'} items-center gap-2 flex-1 h-[100px]`}>
        <RikishiMatchCardImage rikishi={rikishi} />

        <div className="py-2 flex-1 text-center">
          <div className={`font-semibold text-lg sm:text-xl ${language === 'jp' ? 'tracking-widest' : ''}`}>{language === 'en' ? rikishi.shikona_en : rikishi.shikona_jp}</div>
          <div className={`text-sm pb-1 mt-[-0.2rem] text-muted-foreground ${language === 'jp' && 'tracking-widest'}`}>{language === 'en' ? rikishi.banzuke_name_en : rikishi.banzuke_name_jp}</div>
          <div className="text-lg">({rikishi.won} - {rikishi.lost})</div>

          {false && isExpanded && (
            <div className="mt-4 space-y-1 text-[2.5vw] sm:text-[1.8vw] md:text-[1.2vw] text-muted-foreground">
              <div><span className="font-medium">Stable:</span> Miyagino</div>
              <div><span className="font-medium">Name:</span> Yokozuna Hakuho</div>
              <div><span className="font-medium">Ring Name:</span> 白鵬 翔</div>
              <div><span className="font-medium">Birthday:</span> March 11, 1985</div>
              <div><span className="font-medium">Birthplace:</span> Ulaanbaatar, Mongolia</div>
              <div><span className="font-medium">Height:</span> 192 cm</div>
              <div><span className="font-medium">Weight:</span> 155 kg</div>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
