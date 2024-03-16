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

export const repoSchema = z.object({
  id: z.number(),
  name: z.string(),
  connected: z.boolean(),
  creator: z.string(),
  fundedAmount: z.string().optional(),
  description: z.string().optional(),
  amount: z.string().optional(),
  label: z.union([z.string(), z.undefined(), z.null()]),
  status: z.string().optional(),
  url: z.string().optional(),
})

export type IssueTask = z.infer<typeof issueSchema>
export type RepositoryTask = z.infer<typeof repoSchema>
