
import { columns } from "@/components/issues-table/columns"
import { DataTable } from "@/components/issues-table/data-table"
import { UserNav } from "@/components/issues-table/user-nav"
import { Metadata } from "next"
import { z } from "zod"
import { issueSchema } from "../data/schema"

export const metadata: Metadata = {
    title: "Issues Dashboard",
    description: "Gitpaid Issues Dashboard.",
}

async function getIssues() {
    const apiKey = process.env.EMITLY_API_KEY as string
    const url = `https://api.emitly.dev/v1/webhook?listenerId=fn_231fbbe718228828ed3f1d56d88b24e9&apikey=${apiKey}&issues=true`
    const response = await fetch(url, { method: 'POST', cache: 'no-cache' })

    const jsonData = await response.json()
    const outputIssues = []
    for (const issue of jsonData.body) {
        outputIssues.push({
            id: issue.Data.Issue.Number,
            repository: issue.Data.Repo.Name,
            solverAvatar: issue.Metadata?.solverAvatar || issue.Data.Sender.AvatarURL,
            solverUsername: issue.Data.Sender.Login,
            transactionId: issue.Metadata?.txID || "NA",
            description: issue.Data.Issue.Title,
            amount: issue.Metadata?.amount || "NA",
            creator: issue.Data.Sender.Login,
            url: issue.Data.Issue.HTMLURL,
            avatar: issue.Data.Sender.AvatarURL,
            label: issue.Metadata?.label || issue.Data.Issue.Labels.find((label: any) => label)?.Name || "NA",
            status: issue.Data.Issue.State,
            language: issue.Data.Repo.Language || "NA",
        })
    }

    return z.array(issueSchema).parse(outputIssues);
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

interface DataResponse {
    ID: string
    Metadata: IssueMetadata
    Data: IssueObject
}

interface IssueMetadata {
    label: string
    amount: string
    txID: string
    solverAvatar: string
    solverUsername: string
}

interface IssueObject {
    Action: string;
    metadata: IssueMetadata;
    Assignee?: any;
    Changes?: any;
    Installation?: any;
    Issue: Issue;
    Label?: any;
    Repo: Repo;
    Sender: Sender;
}

interface Sender {
    AvatarURL: string;
    Bio?: any;
    Blog?: any;
    Collaborators?: any;
    Company?: any;
    CreatedAt?: any;
    DiskUsage?: any;
    Email?: any;
    EventsURL: string;
    Followers?: any;
    FollowersURL: string;
    Following?: any;
    FollowingURL: string;
    GistsURL: string;
    GravatarID: string;
    HTMLURL: string;
    Hireable?: any;
    ID: number;
    Location?: any;
    Login: string;
    Name?: any;
    NodeID: string;
    OrganizationsURL: string;
    OwnedPrivateRepos?: any;
    Permissions?: any;
    Plan?: any;
    PrivateGists?: any;
    PublicGists?: any;
    PublicRepos?: any;
    ReceivedEventsURL: string;
    ReposURL: string;
    SiteAdmin: boolean;
    StarredURL: string;
    SubscriptionsURL: string;
    SuspendedAt?: any;
    TextMatches?: any;
    TotalPrivateRepos?: any;
    Type: string;
    URL: string;
    UpdatedAt?: any;
}

interface Repo {
    AllowMergeCommit?: any;
    AllowRebaseMerge?: any;
    AllowSquashMerge?: any;
    ArchiveURL: string;
    Archived: boolean;
    AssigneesURL: string;
    AutoInit?: any;
    BlobsURL: string;
    BranchesURL: string;
    CloneURL: string;
    CodeOfConduct?: any;
    CollaboratorsURL: string;
    CommentsURL: string;
    CommitsURL: string;
    CompareURL: string;
    ContentsURL: string;
    ContributorsURL: string;
    CreatedAt: CreatedAt;
    DefaultBranch: string;
    DeploymentsURL: string;
    Description?: any;
    DownloadsURL: string;
    EventsURL: string;
    Fork: boolean;
    ForksCount: number;
    ForksURL: string;
    FullName: string;
    GitCommitsURL: string;
    GitRefsURL: string;
    GitTagsURL: string;
    GitURL: string;
    GitignoreTemplate?: any;
    HTMLURL: string;
    HasDownloads: boolean;
    HasIssues: boolean;
    HasPages: boolean;
    HasProjects: boolean;
    HasWiki: boolean;
    Homepage?: any;
    HooksURL: string;
    ID: number;
    IssueCommentURL: string;
    IssueEventsURL: string;
    IssuesURL: string;
    KeysURL: string;
    LabelsURL: string;
    Language?: any;
    LanguagesURL: string;
    License?: any;
    LicenseTemplate?: any;
    MasterBranch?: any;
    MergesURL: string;
    MilestonesURL: string;
    MirrorURL?: any;
    Name: string;
    NetworkCount?: any;
    NodeID: string;
    NotificationsURL: string;
    OpenIssuesCount: number;
    Organization?: any;
    Owner: any[];
    Parent?: any;
    Permissions?: any;
    Private: boolean;
    PullsURL: string;
    PushedAt: CreatedAt;
    ReleasesURL: string;
    SSHURL: string;
    SVNURL: string;
    Size: number;
    Source?: any;
    StargazersCount: number;
    StargazersURL: string;
    StatusesURL: string;
    SubscribersCount?: any;
    SubscribersURL: string;
    SubscriptionURL: string;
    TagsURL: string;
    TeamID?: any;
    TeamsURL: string;
    TextMatches?: any;
    Topics: any[];
    TreesURL: string;
    URL: string;
    UpdatedAt: CreatedAt;
    WatchersCount: number;
}

interface CreatedAt {
}

interface Issue {
    ActiveLockReason?: any;
    Assignee?: any;
    Assignees: any[];
    Body?: any;
    ClosedAt: string;
    ClosedBy?: any;
    Comments: number;
    CommentsURL: string;
    CreatedAt: string;
    EventsURL: string;
    HTMLURL: string;
    ID: number;
    Labels: any[];
    LabelsURL: string;
    Locked: boolean;
    Milestone?: any;
    NodeID: string;
    Number: number;
    PullRequestLinks?: any;
    Reactions: any[];
    Repository?: any;
    RepositoryURL: string;
    State: string;
    TextMatches?: any;
    Title: string;
    URL: string;
    UpdatedAt: string;
    User: any[];
}