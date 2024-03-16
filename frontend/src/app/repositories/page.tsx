import Table from "@/components/repo-table/table";

function Page({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams?: { [key: string]: string | string[] | undefined };
}) {

    return (
        <main className="pl-24 pr-24 pb-24">
            <Table username={searchParams?.username as string || null} />
        </main>
    )
}

export default Page