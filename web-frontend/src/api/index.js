// 封装 API 请求
const apiBase = '/api'

/**
 * 登录请求
 * @param {string} username - 用户名
 * @param {string} password - 密码
 * @returns {Promise<Response>} 登录请求的响应
 */
export const login = async (username, password) => {
  try {
    const response = await fetch(`${apiBase}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        password
      })
    })
    
    if (!response.ok) {
      throw new Error('登录失败')
    }
    
    return response.json()
  } catch (error) {
    console.error('请求出错:', error)
    throw error
  }
}