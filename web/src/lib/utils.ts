import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

const RIKISHI_NAMES = [
  'Hakuho',
  'Asashoryu',
  'Takanohana',
  'Akebono',
  'Musashimaru',
  'Terunofuji',
  'Takakeisho',
  'Mitakeumi',
  'Konishiki',
  'Chiyonofuji',
  'Wakanohana',
  'Kotomitsuki',
  'Harumafuji',
  'Kakuryu',
  'Kisenosato',
  'Baruto',
  'Goeido',
  'Tochinoshin',
  'Hoshoryu',
  'Kiribayama'
]

export function generateAnonymousRikishiName(): string {
  const randomRikishi = RIKISHI_NAMES[Math.floor(Math.random() * RIKISHI_NAMES.length)]
  const randomNumber = Math.floor(Math.random() * 100) + 1
  return `${randomRikishi}${randomNumber}`
}