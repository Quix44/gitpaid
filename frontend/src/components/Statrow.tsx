import Image from 'next/image'

function Statrow() {
    return (
        <section className="my-8 w-full flex flex-col lg:grid lg:grid-cols-12 gap-5">
            <div className="flex flex-row lg:flex-col lg:grid lg:grid-cols-subgrid items-center bg-[#E9265C] rounded-3xl h-[158px] lg:col-span-4">
                <div className="flex justify-center lg:col-span-2">
                    <Image
                        src="/Nounswizard.png"
                        width={116}
                        height={158}
                        alt="Nouns Wizard"
                    />
                </div>
                <div className="lg:col-span-2">
                    <h4 className="text-xl font-semibold tracking-tight">
                        Total Funds
                    </h4>
                    <h2 className="pb-2 text-3xl font-semibold tracking-tight">
                        $5,423
                    </h2>
                    <p className="leading-7">
                        16% this month
                    </p>
                </div>
            </div>
            <div className="flex flex-row lg:flex-col  lg:grid lg:grid-cols-subgrid items-center bg-[#F1D090] rounded-3xl lg:col-span-4">
                <div className="flex justify-center lg:col-span-2">
                    <Image
                        src="/Nounszebra.png"
                        width={116}
                        height={158}
                        alt="Nouns Zebra"
                    />
                </div>
                <div className="lg:col-span-2">
                    <h4 className="text-xl font-semibold tracking-tight">
                        Active Issues
                    </h4>
                    <h2 className="pb-2 text-3xl font-semibold tracking-tight">
                        1,893
                    </h2>
                    <p className="leading-7">
                        1% this month
                    </p>
                </div>
            </div>
            <div className="flex flex-row lg:flex-col lg:grid lg:grid-cols-subgrid items-center bg-[#34AC80] rounded-3xl lg:col-span-4">
                <div className="flex justify-center lg:col-span-2">
                    <Image
                        src="/Nounsunicorn.png"
                        width={116}
                        height={158}
                        alt="Nouns Unicorn"
                    />
                </div>
                <div className="lg:col-span-2">
                    <h4 className="text-xl font-semibold tracking-tight">
                        Total Paid
                    </h4>
                    <h2 className="pb-2 text-3xl font-semibold tracking-tight">
                        $5,423
                    </h2>
                    <p className="leading-7">
                        1% this month
                    </p>
                </div>
            </div>
        </section>
    )
}

export default Statrow