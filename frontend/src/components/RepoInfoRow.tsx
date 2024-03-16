
function RepoInfoRow() {
    return (
        <section className="grid grid-cols-12 my-3 w-full">
            <div className="grid col-span-4 bg-red-900 grid-cols-subgrid rounded-md items-center h-[158px]">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-[#E9265C] w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">1</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        CONNECT
                    </h2>
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">Your Repositories</h4>
                </div>
            </div>
            <div className="grid col-span-4 grid-cols-subgrid items-center bg-orange-400 rounded-md mx-2">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-[#F1D090] w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">2</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        FUND
                    </h2>
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">With Your Tokens</h4>
                </div>
            </div>
            <div className="grid col-span-4 grid-cols-subgrid items-center bg-lime-700 rounded-md">
                <div className="col-span-2">
                    <div className="grid col-span-2 justify-center items-center rounded-full bg-[#34AC80] w-32 h-32 ml-3"><h1 className=" col-span-5 scroll-m-20 font-extrabold tracking-tight text-8xl items-center">3</h1></div>
                </div>
                <div className="col-span-2 ">
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        APPLY
                    </h2>
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">Your Webhook URL</h4>
                </div>
            </div>

        </section>
    )
}

export default RepoInfoRow