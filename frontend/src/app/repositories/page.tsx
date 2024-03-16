import Table from "@/components/repo-table/table";
import ImportRepo from "../../components/ImportRepo";

function Page({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams?: { [key: string]: string | string[] | undefined };
}) {

    return (
        <main className="p-24">
            <div className="flex items-center justify-between space-y-2">
                <div>
                    <h2 className="text-2xl font-bold tracking-tight title text-primary">Here are your repositories</h2>
                    <p className="text-muted-foreground">
                        Connect, Fund and start creating GitHub Issues!
                    </p>
                </div>
                <ImportRepo />
            </div>
            <Table username={searchParams?.username as string || null} />
        </main>
    )
}

export default Page