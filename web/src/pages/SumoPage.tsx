import { useParams } from 'react-router-dom'
import { NoAuthLayout } from "@/components/layout"
import { MatchList } from "@/containers"
import { useLanguagePreference } from '@/hooks/useLanguagePreference'
import { useKeydownShortcut } from '@/hooks/useKeydownShortcut'
import { useMatchData } from '@/hooks/useMatchData'

export default function SumoPage() {
  const { division, day } = useParams<{ division: string; day: string }>()
  const [language, setLanguage] = useLanguagePreference()
  const { matches, loading, error } = useMatchData(division, day)

  useKeydownShortcut(
    { key: 'l', ctrl: false },
    () => setLanguage(language === 'en' ? 'jp' : 'en'),
    'Toggle Language',
    'Switch between English and Japanese display'
  )

  return (
    <NoAuthLayout className="">
      <div className="space-y-6">
        <div className="text-center">
          <h1 className="text-3xl font-bold text-foreground mb-2">
            Sumo Torikumi
          </h1>
          <p className="text-lg text-muted-foreground">
            Division: {division} | Day: {day}
          </p>
        </div>

        {loading && (
          <div className="text-center text-muted-foreground">
            Loading matches...
          </div>
        )}

        {error && (
          <div className="text-center text-red-500">
            Error: {error}
          </div>
        )}

        {!loading && !error && <MatchList matches={matches} />}
      </div>
    </NoAuthLayout>
  )
}
