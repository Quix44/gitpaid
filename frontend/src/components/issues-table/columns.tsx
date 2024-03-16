"use client"

import { ColumnDef } from "@tanstack/react-table"


import Image from "next/image"
import { statuses } from "../data/data"
import { IssueTask } from "../data/schema"
import { Badge } from "../ui/badge"
import {
  Tooltip
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

      if (!status) {
        return null
      }
      const textColorClass = status ? 'text-green-500' : 'text-red-500';

      return (
        <div className="flex w-[100px] items-center">
          {status && status.icon && (
            <status.icon className={`mr-2 h-4 w-4 ${textColorClass}`} />
          )}
          {!status ? (
            // When there's no status, wrap the "No Status" text in a Tooltip
            <Tooltip content="No Status Available" position="top">
              <span className={textColorClass}>No Status</span>
            </Tooltip>
          ) : (
            // Render status label with appropriate text color when status exists
            <span className={textColorClass}>{status.label}</span>
          )}
        </div>
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
