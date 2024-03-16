import Image from 'next/image'

function Statrow() {
    return (
        <section className="grid grid-cols-12 my-8 w-full">
            <div className="grid col-span-4 bg-red-900 grid-cols-subgrid rounded-md items-center h-[158px]">
                <div className="col-span-2">
                    <Image
                        src="/Nounswizard.png"
                        width={116}
                        height={158}
                        alt="Nouns Wizard"
                    />
                </div>
                <div className="col-span-2 ">
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        Total Funds
                    </h4>
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        $5,423
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        16% this month
                    </p>
                </div>
            </div>
            <div className="grid col-span-4 grid-cols-subgrid items-center bg-orange-400 rounded-md mx-2">
                <div className="col-span-2">
                    <Image
                        src="/Nounszebra.png"
                        width={116}
                        height={158}
                        alt="Nouns Zebra"
                    />
                </div>
                <div className="col-span-2">
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        Active Issues
                    </h4>
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        1,893
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        1% this month
                    </p>
                </div>
            </div>
            <div className="grid col-span-4 grid-cols-subgrid items-center bg-lime-700 rounded-md">
                <div className="col-span-2">
                    <Image
                        src="/Nounsunicorn.png"
                        width={116}
                        height={158}
                        alt="Nouns Unicorn"
                    />
                </div>
                <div className="col-span-2">
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        Total Paid
                    </h4>
                    <h2 className="scroll-m-20  pb-2 text-3xl font-semibold tracking-tight first:mt-0">
                        $5,423
                    </h2>
                    <p className="leading-7 [&:not(:first-child)]:mt-0">
                        1% this month
                    </p>
                </div>
            </div>

        </section>
    )
}

export default Statrow