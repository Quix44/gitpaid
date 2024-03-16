
function Searchinfo() {
    return (
        <section className="pt-6">
            <div className="grid grid-cols-12 w-full">
                <div className="grid col-span-3">
                    <h1 className="grid col-span-2 scroll-m-20 font-extrabold tracking-tight lg:text-5xl items-center">
                        Git Started
                    </h1>
                </div>

                <div className="grid col-start-6 col-span-7 grid-cols-subgrid space-y-12 items-center">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-[#E9265C] w-32 h-32 mt-[48px]">
                        <h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            1
                        </h1>
                    </div>
                    <div className="col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Lets log in
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            Simply login with your github profile to get started
                        </p>
                    </div>

                    <div className="grid col-span-2 justify-center items-center justify-self rounded-full w-32 h-32 bg-[#F1D090]">
                        <h1 className="col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            2
                        </h1>
                    </div>
                    <div className="col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Search and commit
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            You can search through our global issue repo based on price, skill and type.
                        </p>
                    </div>

                    <div className="grid col-span-2 justify-center items-center justify-self rounded-full w-32 h-32 bg-[#34AC80]">
                        <h1 className="col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">
                            3
                        </h1>
                    </div>
                    <div className="col-span-5">
                        <div className="text-lg font-semibold text-[#46978E]">
                            Search and commit
                        </div>
                        <p className="leading-7 [&:not(:first-child)]:mt-0">
                            Submit a PR to your chosen issue, if the PR is accepted, you’ll be getting those juicy dollar bills in your wallet.
                        </p>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Searchinfo