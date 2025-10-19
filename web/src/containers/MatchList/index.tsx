import { useParams } from 'react-router-dom'
import { RikishiMatchCard } from '@/components/ui'
import { matches } from '@/stubs/sumo'
import { useKeydownShortcut } from '@/hooks/useKeydownShortcut'
import { useLanguagePreference } from '@/hooks/useLanguagePreference'

export default function MatchList() {
  const { division, day } = useParams<{ division: string; day: string }>()
  const [language, setLanguage] = useLanguagePreference()

  useKeydownShortcut(
    { key: 'l', ctrl: false },
    () => setLanguage(prev => prev === 'en' ? 'jp' : 'en'),
    'Toggle Language',
    'Switch between English and Japanese display'
  )

  return (
    <div className="space-y-4">
      {/* Column headers */}
      <div className="flex gap-4 text-center font-semibold text-muted-foreground">
        <div className="flex-4 md:flex-2">East</div>
        <div className="flex-1"></div>
        <div className="flex-4 md:flex-2">West</div>
      </div>

      {/* Match rows */}
      {matches.map((match, index) => (
        <div key={index} className="flex gap-2 items-start border border-border shadow-sm rounded md:h-[100px] overflow-hidden">
          <div className="flex-4 md:flex-2 h-full">
            <RikishiMatchCard rikishi={match.east} side="east" language={language} />
          </div>
          <div className="flex h-full"></div>
          <div className="flex-4 md:flex-2 h-full">
            <RikishiMatchCard rikishi={match.west} side="west" language={language} />
          </div>
        </div>
      ))}
    </div>
  )
}
