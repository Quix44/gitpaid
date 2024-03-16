import ImportRepo from "../../components/ImportRepo"

function page() {
    return (
        <main className="p-24">
            <div className="flex items-center justify-between space-y-2">
                <div>
                    <h2 className="text-2xl font-bold tracking-tight title text-primary">Here are your repositories</h2>
                    <p className="text-muted-foreground">
                        Here&apos;s a list of your tasks for this month!
                    </p>
                </div>
                <ImportRepo />
            </div>
        </main>
    )
}

export default page