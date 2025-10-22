import { useLanguagePreference } from '@/hooks/useLanguagePreference'

export default function RikishiMatchListColumns() {
  const [language] = useLanguagePreference()
  const eastText = language === 'en' ? 'East' : '東'
  const westText = language === 'en' ? 'West' : '西'

  return (
    <div className="flex gap-4 font-semibold text-muted-foreground">
      <div className="flex-4 text-center text-xl pl-2">{eastText}</div>
      <div className="flex h-full"></div>
      <div className="flex-4 text-center text-xl pr-2">{westText}</div>
    </div>
  )
}
