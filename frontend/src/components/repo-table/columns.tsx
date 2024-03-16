"use client"

import Provider from "@/app/(providers)/Wallet"
import { ColumnDef } from "@tanstack/react-table"
import ConnectRepositoryButton from "../ConnectRepositoryButton"
import { RepositoryTask } from "../data/schema"
import { FundButton } from "../fund-button"
import { DataTableColumnHeader } from "./data-table-column-header"
import { DataTableRowActions } from "./data-table-row-actions"

const openRepo = (repoURL: string | undefined) => {
  if (!repoURL) return
  window.open(repoURL, '_blank', 'noopener,noreferrer');
}
export const columns: ColumnDef<RepositoryTask>[] = [
  {
    accessorKey: "name",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Name" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label)

      return (
        <div className="flex space-x-2 max-w-[5rem]">
          {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
          <span className="max-w-[500px] truncate font-medium">
            {row.getValue("name")}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "description",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Description" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label)

      return (
        <div className="flex space-x-2 cursor-pointer" onClick={() => openRepo(row.original.url)}>
          {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
          <span className="max-w-[200px] truncate font-medium">
            {row.getValue("description")}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "fundedAmount",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Funded Amount" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label)

      return (
        <div className="flex space-x-2 cursor-pointer" onClick={() => openRepo(row.original.fundedAmount)}>
          {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
          <span className="max-w-[200px] truncate font-medium">
            {row.getValue("fundedAmount")}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "creator",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Creator" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label)

      return (
        <div className="flex space-x-2">
          {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
          <span className="max-w-[500px] truncate font-medium">
            {row.getValue("creator")}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "label",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Label" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label)

      return (
        <div className="flex space-x-2">
          {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
          <span className="max-w-[500px] truncate font-medium">
            {row.getValue("label")}
          </span>
        </div>
      )
    },
  },
  // {
  //   accessorKey: "status",
  //   header: ({ column }) => (
  //     <DataTableColumnHeader column={column} title="Status" />
  //   ),
  //   cell: ({ row }) => {
  //     const status = statuses.find(
  //       (status) => status.value === row.getValue("status")
  //     )

  //     if (!status) {
  //       return null
  //     }

  //     return (
  //       <div className="flex w-[100px] items-center">
  //         {status.icon && (
  //           <status.icon className="mr-2 h-4 w-4 text-muted-foreground" />
  //         )}
  //         <span>{status.label}</span>
  //       </div>
  //     )
  //   },
  //   filterFn: (row, id, value) => {
  //     return value.includes(row.getValue(id))
  //   },
  // },
  {
    accessorKey: "connect",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Connect" />
    ),
    cell: ({ row }) => {

      return (
        <ConnectRepositoryButton connected={row.original.connected} />
      )
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id))
    },
  },
  {
    accessorKey: "fund",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Fund" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex w-[100px] items-center">
          <Provider>
            <FundButton repository={String(row.original.id)} />
          </Provider>
        </div>
      )
    },
  },
  {
    id: "actions",
    cell: ({ row }) => <DataTableRowActions row={row} />,
  },
]


