import { CreateTasks } from "@/types/Tasks";
import axiosInstance from "@/utils/httpClient";
import { AxiosResponse } from "axios";
const PATH = "tasks";
const findAll = async (params = {}): Promise<AxiosResponse> => {
  return await axiosInstance.get(`/${PATH}`, { params });
};

const findById = async (id: string): Promise<AxiosResponse> => {
  return await axiosInstance.get(`/${PATH}/${id}`);
};

const create = async (data: CreateTasks): Promise<AxiosResponse> => {
  return await axiosInstance.post(PATH, data);
};

const update = async (
  id: string,
  data: CreateTasks
): Promise<AxiosResponse> => {
  return await axiosInstance.put(`/${PATH}/${id}`, data);
};

const remove = async (id: string): Promise<AxiosResponse> => {
  return await axiosInstance.delete(`/${PATH}/${id}`);
};

const TaskService = {
  findAll,
  findById,
  create,
  update,
  remove,
};

export default TaskService;
