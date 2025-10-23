import { useState, useRef, useEffect } from 'react'
import { RikishiMatchCard, RikishiMatchListColumns } from '@/components/ui'
import type { Match } from '@/types'
import { MATCH_SIDE } from '@/types'

interface MatchListProps {
  matches: Match[]
}

export default function MatchList({ matches }: MatchListProps) {
  const [expandedRow, setExpandedRow] = useState<number | null>(null)
  const rowRefs = useRef<(HTMLDivElement | null)[]>([])

  useEffect(() => {
    if (expandedRow !== null && rowRefs.current[expandedRow]) {
      const element = rowRefs.current[expandedRow]
      if (element) {
        const elementPosition = element.getBoundingClientRect().top + window.scrollY
        const offsetPosition = elementPosition - 80

        window.scrollTo({
          top: offsetPosition,
          behavior: 'smooth'
        })
      }
    }
  }, [expandedRow])

  return (
    <div className="m-auto sm:max-w-2xl space-y-4">
      {/* Column headers */}
      <RikishiMatchListColumns />

      {/* Match rows */}
      {matches.map((match, index) => (
        <div
          ref={(el) => (rowRefs.current[index] = el)}
          key={index}
          onClick={() => setExpandedRow(expandedRow === index ? null : index)}
          className={`flex gap-2 items-start border border-border shadow-sm rounded overflow-hidden transition-all duration-300 cursor-pointer hover:shadow-md`}
        >
          <div className="flex-4 md:flex-2 h-full">
            <RikishiMatchCard rikishi={match.east} side={MATCH_SIDE.EAST} />
          </div>
          <div className="flex h-full"></div>
          <div className="flex-4 md:flex-2 h-full">
            <RikishiMatchCard rikishi={match.west} side={MATCH_SIDE.WEST} />
          </div>
        </div>
      ))}
    </div>
  )
}
