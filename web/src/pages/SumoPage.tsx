import { useParams } from 'react-router-dom'
import { NoAuthLayout } from "@/components/layout"
import { MatchList } from "@/containers"
import { matches } from '@/stubs/sumo'
import { useLanguagePreference } from '@/hooks/useLanguagePreference'
import { useKeydownShortcut } from '@/hooks/useKeydownShortcut'

export default function SumoPage() {
  const { division, day } = useParams<{ division: string; day: string }>()
  const [language, setLanguage] = useLanguagePreference()

  useKeydownShortcut(
    { key: 'l', ctrl: false },
    () => setLanguage(language === 'en' ? 'jp' : 'en'),
    'Toggle Language',
    'Switch between English and Japanese display'
  )

  return (
    <NoAuthLayout>
      <div className="space-y-6">
        <div className="text-center">
          <h1 className="text-3xl font-bold text-foreground mb-2">
            Sumo Torikumi
          </h1>
          <p className="text-lg text-muted-foreground">
            Division: {division} | Day: {day}
          </p>
        </div>

        <MatchList matches={matches} />
      </div>
    </NoAuthLayout>
  )
}
