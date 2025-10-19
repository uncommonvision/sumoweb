import type { Language, Rikishi } from '@/types'

interface RikishiMatchCardProps {
  rikishi: Rikishi
  side: 'east' | 'west'
  language: Language
}

export default function RikishiMatchCard({ rikishi, side, language }: RikishiMatchCardProps) {
  const isEast = side === 'east'
  return (
    <div className="bg-card text-card-foreground flex items-center h-full">
      <div className={`flex ${isEast ? 'flex-row' : 'flex-row-reverse'} items-center gap-2 flex-1 h-full`}>
        <div className="md:flex hidden overflow-hidden self-start">
          <img
            src="https://www.sumo.or.jp/img/sumo_data/rikishi/270x474/20150005.jpg"
            alt="Rikishi"
            className="h-auto max-w-[200px] object-cover object-top scale-125 origin-top rounded"
          />
        </div>
        <div className="py-2 flex-1 text-center">
          <div className={`font-semibold text-[4vw] sm:text-[3vw] md:text-[2vw] ${language === 'jp' ? 'tracking-widest' : ''}`}>{language === 'en' ? rikishi.shikona_en : rikishi.shikona_jp}</div>
          <div className={`text-[3vw] sm:text-[2vw] md:text-[1.5vw] text-muted-foreground ${language === 'jp' ? 'tracking-widest' : ''}`}>{language === 'en' ? rikishi.banzuke_name_en : rikishi.banzuke_name_jp}</div>
          <div className="text-[3vw] sm:text-[2vw] md:text-[1.5vw]">({rikishi.won} - {rikishi.lost})</div>
        </div>
        <div className="bg-gray-200 rounded p-1 md:p-2">
          <div className={`w-4 h-4 rounded-full shadow ${rikishi.result ? 'bg-white' : 'bg-black'}`}></div>
        </div>
      </div>
    </div>
  )
}
