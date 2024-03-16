import Provider from "@/app/(providers)/Wallet"
import { columns } from "@/components/repo-table/columns"
import { ethers } from 'ethers'
import { Metadata } from "next"
import { z } from "zod"
import ImportRepo from "../ImportRepo"
import RepoInfoRow from "../RepoInfoRow"
import { repoSchema } from "../data/schema"
import { DataTable } from "./data-table"

export const metadata: Metadata = {
    title: "Repositories Dashboard",
    description: "Gitpaid Repositories Dashboard.",
}

async function getRepos(username: string | null) {
    if (!username) return []
    const apiKey = process.env.EMITLY_API_KEY as string
    const url = `https://api.emitly.dev/v1/webhook?listenerId=fn_231fbbe718228828ed3f1d56d88b24e9&apikey=${apiKey}&repos=true&user=${username}`
    const response = await fetch(url, { method: 'POST', cache: 'no-cache' })
    const jsonData = await response.json()

    const outputRepositories = []
    for (const repo of jsonData.body) {
        outputRepositories.push({
            id: repo.Data.id,
            name: repo.Data.name,
            description: repo.Data.description || "NA",
            connected: repo.Metadata ? true : false,
            fundedAmount: Math.floor(+ethers.utils.formatUnits(repo.Metadata?.amount || "0", 'ether')).toString(),
            creator: repo.Data.owner.login,
            label: repo.Data.language,
            amount: repo.Metadata?.amount ?? "",
            url: repo.Data.html_url ?? ""
        })
    }

    return z.array(repoSchema).parse(outputRepositories);
}

export async function Table({ username }: { username: string | null }) {
    const tasks = await getRepos(username)
    return (
        <div className={`hidden h-full flex-1 flex-col space-y-8 p-8 md:flex`}>
            <Provider>
                <RepoInfoRow />
                <div className="flex justify-between w-full">
                    <div>
                        <h2 className="text-2xl font-bold tracking-tight title text-white">Here are your repositories</h2>
                        <p className="text-muted-foreground">
                            Connect, Fund and start creating GitHub Issues!
                        </p>
                    </div>
                    <ImportRepo />
                </div>
                <DataTable data={tasks} columns={columns} />
            </Provider>
        </div >
    )
}
export default Table

interface DataResponse {
    ID: string
    Metadata: RepoMetadata
    Data: Data
}

interface Data {
    id: number;
    allow_forking: boolean;
    archived: boolean;
    archive_url: string;
    assignees_url: string;
    blobs_url: string;
    branches_url: string;
    clone_url: string;
    collaborators_url: string;
    comments_url: string;
    commits_url: string;
    compare_url: string;
    contents_url: string;
    contributors_url: string;
    created_at: string;
    default_branch: string;
    deployments_url: string;
    description?: any;
    disabled: boolean;
    downloads_url: string;
    events_url: string;
    fork: boolean;
    forks: number;
    forks_count: number;
    forks_url: string;
    full_name: string;
    git_commits_url: string;
    git_refs_url: string;
    git_tags_url: string;
    git_url: string;
    has_discussions: boolean;
    has_downloads: boolean;
    has_issues: boolean;
    has_pages: boolean;
    has_projects: boolean;
    has_wiki: boolean;
    homepage?: any;
    hooks_url: string;
    html_url: string;
    issues_url: string;
    issue_comment_url: string;
    issue_events_url: string;
    is_template: boolean;
    keys_url: string;
    labels_url: string;
    language: string;
    languages_url: string;
    license?: any;
    merges_url: string;
    milestones_url: string;
    mirror_url?: any;
    name: string;
    node_id: string;
    notifications_url: string;
    open_issues: number;
    open_issues_count: number;
    owner: Owner;
    private: boolean;
    pulls_url: string;
    pushed_at: string;
    releases_url: string;
    size: number;
    ssh_url: string;
    stargazers_count: number;
    stargazers_url: string;
    statuses_url: string;
    subscribers_url: string;
    subscription_url: string;
    svn_url: string;
    tags_url: string;
    teams_url: string;
    topics: any[];
    trees_url: string;
    updated_at: string;
    url: string;
    visibility: string;
    watchers: number;
    watchers_count: number;
    web_commit_signoff_required: boolean;
}

interface Owner {
    id: number;
    avatar_url: string;
    events_url: string;
    followers_url: string;
    following_url: string;
    gists_url: string;
    gravatar_id: string;
    html_url: string;
    login: string;
    node_id: string;
    organizations_url: string;
    received_events_url: string;
    repos_url: string;
    site_admin: boolean;
    starred_url: string;
    subscriptions_url: string;
    type: string;
    url: string;
}
interface RepoMetadata {
    amount: string;
    chainID: string;
    contractAddress: string;
    payeeAddress: string;
    tokenSymbol: string;
    tokenSymbolUrl?: string;
    rpc: string;
    tokenAddress: string;
}