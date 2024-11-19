import TaskService from "@/services/TaskService";
import { defineStore } from "pinia";
import { Response } from "@/types/Response";
import { toast } from "vue3-toastify";
import { Tasks } from "@/types/Tasks";
import _ from "lodash";

interface TaskStateType extends Response {
  loading: boolean;
}

export const useTaskStore = defineStore("tasks", () => {
  const statusList = [
    {
      label: "ฉบับร่าง",
      value: "draft",
    },
    {
      label: "กำลังดําเนินการ",
      value: "in_progress",
    },
    {
      label: "ดําเนินการเสร็จสิ้น",
      value: "completed",
    },
  ];

  const initStateTask: TaskStateType = {
    error: false,
    message: "",
    result: [],
    status: 0,
    statusText: "",
    loading: false,
  };

  let task = reactive<Tasks>({
    id: "",
    title: "",
    description: "",
    status: "",
    created_at: new Date(),
    updated_at: new Date(),
  });

  let tasks = reactive<TaskStateType>(initStateTask);

  let search = reactive({
    search: "",
    status: null,
  });

  const findAllTasks = async () => {
    tasks.loading = true;

    try {
      const {
        data,
      }: {
        data: Response;
        status: number;
      } = await TaskService.findAll(search);
      // ใช้ Object.assign เพื่ออัพเดท  tasks
      Object.assign(tasks, initStateTask, data);
      return tasks;
    } catch (error) {
      console.error("findAllTasks error", error);
    } finally {
      tasks.loading = false;
    }
  };

  const deleteTaskById = async (id: string) => {
    tasks.loading = true;
    try {
      const {
        data,
        status,
      }: {
        data: Response;
        status: number;
      } = await TaskService.remove(id);
      const { error, message } = data;

      // ลบไม่สําเร็จ
      if (status !== 200 || error) {
        toast.error("ลบข้อมูลไม่สําเร็จ");
        console.error("deleteTaskById error", message);
        const response = data;
        delete response.result;

        // ใช้ Object.assign เพื่ออัพเดท  tasks
        Object.assign(tasks, initStateTask, response);

        return false;
      }

      const result = tasks.result.filter((item: any) => item.id !== id);
      tasks.result = result;
      toast.success("ลบข้อมูลเรียบร้อย");
      return tasks;
    } catch (error) {
      console.error("deleteTaskById error", error);
      toast.error("ลบข้อมูลไม่สําเร็จ");
    } finally {
      tasks.loading = false;
    }
  };

  return {
    // export state
    tasks,
    task,
    statusList,
    search,
    // export actions
    findAllTasks,
    deleteTaskById,
  };
});
