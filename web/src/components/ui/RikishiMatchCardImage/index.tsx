import type { Rikishi } from '@/types'
import { useLanguagePreference } from '@/hooks/useLanguagePreference'

interface RikishiMatchCardImageProps {
  rikishi: Rikishi
}

export default function RikishiMatchCardImage({ rikishi }: RikishiMatchCardImageProps) {
  const [language] = useLanguagePreference()

  if (!rikishi.profile) {
    return (
      <div className="hidden sm:flex h-[100px] sm:w-[6rem] md:w-[8rem] bg-top bg-center bg-clip-content">
      </div>
    )
  }

  // encountering a bug on the english profile page
  // using the japanese image until a fix can be found
  // const imageSrc = window.location.protocol + "//" + window.location.host + rikishi.profile.profiles[language].image_url
  const imageSrc = window.location.protocol + "//" + window.location.host + rikishi.profile.profiles["jp"].image_url

  return (
    <div
      className={`hidden sm:flex h-[100px] sm:w-[6rem] md:w-[8rem] bg-top bg-center bg-size-[125px] bg-clip-content`}
      style={{ backgroundImage: `url(${imageSrc})` }}
    >
    </div>
  )
}
