import type {Task} from "@/models/tasks";

export interface TasksListRequest {}

export interface TasksListResponse {
  totalCount: number,
  items: Task[]
}
