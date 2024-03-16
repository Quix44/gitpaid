import Table from "@/components/repo-table/table";

function Page({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams?: { [key: string]: string | string[] | undefined };
}) {

    return (
        <main className="p-24">
            <Table username={searchParams?.username as string || null} />
        </main>
    )
}

export default Page