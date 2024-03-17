
function RepoInfoRow() {
    return (
        <section className="grid grid-cols-12 my-3 w-full space-x-5">
            <div className="grid col-span-4 rounded-lg border backdrop-blur-xl bg-card text-card-foreground shadow-sm grid-cols-subgrid items-center h-[158px] ">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-gradient w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">1</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        CONNECT
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        Your Repositories
                    </p>
                </div>
            </div>

            <div className="grid col-span-4 grid-cols-subgrid items-center rounded-lg border backdrop-blur-xl bg-card text-card-foreground shadow-sm mx-2">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-gradient w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">2</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        LINK
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        Your Webhook URL
                    </p>
                </div>
            </div>
            <div className="grid col-span-4 grid-cols-subgrid items-center rounded-lg border backdrop-blur-xl bg-card text-card-foreground shadow-sm">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-gradient w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">3</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        FUND
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        With any ERC20 Token
                    </p>

                </div>
            </div>

        </section>
    )
}

export default RepoInfoRow