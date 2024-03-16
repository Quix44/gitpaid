
function Searchinfo() {
    return (
        <section className="my-12">
            <div className="grid grid-cols-12 w-full">
                <div className="col-span-12 lg:col-span-3">
                    <h1 className="text-center lg:text-left text-5xl font-extrabold tracking-tight lg:grid items-center">
                        Git Started
                    </h1>
                </div>

                <div className="grid col-span-12 lg:col-start-6 lg:col-span-7 grid-cols-subgrid space-y-12 items-center">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-[#E9265C] w-32 h-32 mt-[48px]">
                        <h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            1
                        </h1>
                    </div>
                    <div className="col-span-12 lg:col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Lets log in
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            Simply login with your Github account to get started
                        </p>
                    </div>

                    <div className="grid col-span-12 lg:col-span-2 justify-center items-center justify-self rounded-full w-32 h-32 bg-[#F1D090]">
                        <h1 className="col-span-12 lg:col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            2
                        </h1>
                    </div>
                    <div className="col-span-12 lg:col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Search and commit
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            Search through our aggregated set of incentivized issues based on price, language and type.
                        </p>
                    </div>

                    <div className="grid col-span-12 lg:col-span-2 justify-center items-center justify-self rounded-full w-32 h-32 bg-[#34AC80]">
                        <h1 className="col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            3
                        </h1>
                    </div>
                    <div className="col-span-12 lg:col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Search and commit
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            Submit a PR to any issue, if the PR is accepted, you will instantly receive the ERC20 token defined in the issue label.
                        </p>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Searchinfo