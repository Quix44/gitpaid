import Table from "@/components/repo-table/table";

function Page({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams?: { [key: string]: string | string[] | undefined };
}) {

    return (
        <main className="pb-24 pl-24 pr-24">
            <Table username={searchParams?.username as string || null} />
        </main>
    )
}

export default Page