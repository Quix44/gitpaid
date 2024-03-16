import { z } from "zod"

// We're keeping a simple non-relational schema here.
// IRL, you will have a schema for your data models.
export const issueSchema = z.object({
  id: z.number(),
  repository: z.string(),
  description: z.string(),
  creator: z.string(),
  label: z.string(),
  status: z.string(),
  avatar: z.string().optional(),
  language: z.string().optional(),
  url: z.string(),
})

export type Task = z.infer<typeof issueSchema>
