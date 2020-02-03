import axios from 'axios'

// 创建视频
const postVideo = form => axios.post('/api/v1/video', form).then(res => res.data)

// 获取视频
const getVideo = id => axios.get(`/api/v1/video/${id}`, id).then(res => res.data)

// 获取视频列表
const getVideos = form => axios.get('/api/v1/videos', form).then(res => res.data)

export {
  postVideo,
  getVideo,
  getVideos
}
