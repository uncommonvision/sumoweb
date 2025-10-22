import type { Rikishi } from '@/types'

interface RikishiMatchCardImageProps {
  rikishi: Rikishi
}

export default function RikishiMatchCardImage({ rikishi }: RikishiMatchCardImageProps) {
  return (
    <div className={`hidden sm:flex h-[100px] sm:w-[6rem] md:w-[8rem] bg-top bg-center bg-clip-content bg-[url(https://www.sumo.or.jp/img/sumo_data/rikishi/270x474/20150005.jpg)]`}>
    </div>
  )
}
