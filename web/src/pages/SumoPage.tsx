import { useParams } from 'react-router-dom'
import { DefaultLayout } from "@/components/layout"
import { MatchList } from "@/containers"

export default function SumoPage() {
  const { division, day } = useParams<{ division: string; day: string }>()

  return (
    <DefaultLayout>
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-bold text-foreground mb-2">
            Sumo Torikumi
          </h1>
          <p className="text-lg text-muted-foreground">
            Division: {division} | Day: {day}
          </p>
        </div>
        <MatchList />
      </div>
    </DefaultLayout>
  )
}