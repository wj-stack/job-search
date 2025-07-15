export default [
  {
    url: '/api/login',
    method: 'post',
    response: () => {
      return {
        code: 200,
        message: '登录成功',
        data: {
          token: 'mock_token'
        }
      }
    }
  }
]