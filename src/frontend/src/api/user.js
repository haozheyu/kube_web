import {post,get} from "@/plugin/utils/request";

export const login = (params) => post('/login', params)

export const currentuser = (params) => get('/currentuser', params)

