
import { columns } from "@/components/issues-table/columns"
import { DataTable } from "@/components/issues-table/data-table"
import { UserNav } from "@/components/issues-table/user-nav"
import { Metadata } from "next"

export const metadata: Metadata = {
    title: "Tasks",
    description: "A task and issue tracker build using Tanstack Table.",
}

// Simulate a database read for tasks.
async function getIssues() {
    const apiKey = process.env.EMITLY_API_KEY as string
    console.log(apiKey)
    const url = `https://api.emitly.dev/v1/webhook?listenerId=fn_231fbbe718228828ed3f1d56d88b24e9&apikey=${apiKey}&issues=true`

    const response = await fetch(url)
    console.log(response)
    const jsonData = await response.json()
    console.log(jsonData)
    // Get Issues from DyanmoDB
    return jsonData
}

export async function Table() {
    const tasks = await getIssues()

    return (
        <div className={`hidden h-full flex-1 flex-col space-y-8 p-8 md:flex`}>
            <div className="flex items-center justify-between space-y-2">
                <div>
                    <h2 className="text-2xl font-bold tracking-tight">Open Issues</h2>
                    <p className="text-muted-foreground">
                        Here&apos;s the list of open issues for you to work on.
                    </p>
                </div>
                <div className="flex items-center space-x-2">
                    <UserNav />
                </div>
            </div>
            <DataTable data={tasks} columns={columns} />
        </div>
    )
}

export default Table