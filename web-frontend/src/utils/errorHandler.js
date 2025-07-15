import { ElMessage } from 'element-plus'

// 统一错误处理
const showError = (message) => {
  ElMessage.error(message)
}

const handleApiError = (error) => {
  let errorMessage = '请求发生错误，请稍后重试'
  
  if (error.message) {
    errorMessage = error.message
  }
  
  showError(errorMessage)
  throw error
}

 export { handleApiError, showError }