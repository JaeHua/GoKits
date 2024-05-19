import request from '@/utils/request'
import { ElMessage } from 'element-plus'

//发送todo
const postTodo = ({title,status,telephone}) =>{

    return request.post('auth/todo',{title,status,telephone}).then(response => {
        // 请求成功
        ElMessage.success('保存成功')

        return response
    })
        .catch(error => {
            // 请求失败
            ElMessage.error('保存失败')
            return Promise.reject(error)
        })
}
//获取todo
const getTodo = () =>{
    return request.get('auth/todo')
}

//删除todo
const deleteTodo = (id) =>{
    return request.delete('auth/todo:'+id).then(response => {
        // 请求成功
        ElMessage.success('删除成功')

        return response
    })
        .catch(error => {
            // 请求失败
            ElMessage.error('删除失败')
            return Promise.reject(error)
        })
}
//更新todo
const updateTodo = (id) =>{
    return request.put('auth/todo:'+id).then(response => {
        // 请求成功
        ElMessage.success('更新成功')

        return response
    })
        .catch(error => {
            // 请求失败
            ElMessage.error('更新失败')
            return Promise.reject(error)
        })
}
export default {
    postTodo,
    getTodo,
    deleteTodo,
    updateTodo,
}