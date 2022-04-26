export interface Task {
  id?: number,
  title: string,
  description?: string,
  status: string
}

export const statuses = [
  {
    title: 'Todo',
    value: 'todo'
  },
  {
    title: 'In Progress',
    value: 'inProgress',
  },
  {
    title: 'Closed',
    value: 'closed'
  }
]
