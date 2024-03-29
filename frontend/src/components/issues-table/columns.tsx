"use client"

import { ColumnDef } from "@tanstack/react-table"


import Image from "next/image"
import { statuses } from "../data/data"
import { IssueTask } from "../data/schema"
import { Badge } from "../ui/badge"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "../ui/tooltip"
import { DataTableColumnHeader } from "./data-table-column-header"
import { DataTableRowActions } from "./data-table-row-actions"

export const columns: ColumnDef<IssueTask>[] = [
  {
    accessorKey: "id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Task" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("id")}</div>,
    enableSorting: true,
    enableHiding: false,
  },
  {
    accessorKey: "repository",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Repository" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium cursor-pointer hover:text-primary" onClick={() => { window.open(row.original.url, '_blank', 'noreferrer') }}>
            {row.getValue("repository")}
          </span>
        </div >
      )
    },
  },
  {
    accessorKey: "description",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Description" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium">
            {row.getValue("description")}
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
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium flex">
            {row.original.avatar && <Image src={row.original.avatar} className="rounded-full" alt={row.original.creator} width={34} height={34} />}
            {<p className="mt-2 ml-1">{row.getValue("creator")}</p>}
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
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium">
            {row.original.label && <Badge variant="outline">{row.original.label}</Badge>}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "language",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Language" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium">
            {row.original.language && <Badge variant="outline">{row.original.language}</Badge>}
          </span>
        </div>
      )
    },
  },
  {
    accessorKey: "status",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Status" />
    ),
    cell: ({ row }) => {
      const status = statuses.find(
        (status) => status.value === row.getValue("status")
      )

      const isStatusPresent = status?.value === 'open'
      const textColorClass = isStatusPresent ? 'text-green-500' : 'text-red-500'; // Choose color based on status presence

      if (!status) {
        return null
      }

      return (
        <div className="flex w-[100px] items-center">
          {isStatusPresent && status.icon && (
            <status.icon className={`mr-2 h-4 w-4 ${textColorClass}`} />
          )}
          {isStatusPresent ? (
            // If status is present, display its label with the appropriate text color
            <span className={textColorClass}>{status.label}</span>
          ) : (
            // If status is not present, show "No Status" with a tooltip in red
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <span className="text-yellow-500">Paid</span>
                </TooltipTrigger>
                <TooltipContent className="bg-card max-w-sm space-y-2">
                  <h4 className="font-medium leading-none">Issue resolved</h4>
                  <p className="text-sm text-muted-foreground">
                    amount: <b>{row.original.amount}</b>
                  </p>
                  <p className="text-sm text-muted-foreground line-clamp-2 max-w[60px] text-ellipsis ">
                    resolved by: <b>{row.original.solverUsername}</b>
                  </p>
                  <p className="text-sm text-muted-foreground line-clamp-2 max-w[60px] text-ellipsis ">
                    TxID: <b><a href={"https://sepolia.arbiscan.io/tx/" + row.original.transactionId}>{row.original.transactionId}</a></b>
                  </p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          )
          }
        </div >
      )
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id))
    },
  },
  {
    id: "actions",
    cell: ({ row }) => <DataTableRowActions row={row} />,
  },
]
