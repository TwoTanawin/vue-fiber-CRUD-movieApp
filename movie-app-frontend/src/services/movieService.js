import axios from 'axios';

const API_URL = 'http://localhost:3000/movies';

export const getMovies = () => axios.get(API_URL);
export const getMovie = (id) => axios.get(`${API_URL}/${id}`);
export const createMovie = (movie) => axios.post(API_URL, movie);
export const updateMovie = (id, movie) => axios.put(`${API_URL}/${id}`, movie);
export const deleteMovie = (id) => axios.delete(`${API_URL}/${id}`);
