<template>
  <div id="app">
    <h1>Movie List</h1>
    <form @submit.prevent="addMovie">
      <input v-model="title" placeholder="Title" />
      <input v-model="director" placeholder="Director" />
      <input v-model="year" placeholder="Year" type="number" />
      <button type="submit">Add Movie</button>
    </form>
    <ul>
      <li v-for="movie in movies" :key="movie.id">
        {{ movie.title }} by {{ movie.director }} ({{ movie.year }})
        <button @click="editMovie(movie)">Edit</button>
        <button @click="deleteMovie(movie.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { getMovies, createMovie, deleteMovie } from './services/movieService';

export default {
  data() {
    return {
      movies: [],
      title: '',
      director: '',
      year: null,
      selectedMovie: null,
    };
  },
  methods: {
    async fetchMovies() {
      const response = await getMovies();
      this.movies = response.data;
    },
    async addMovie() {
      if (this.selectedMovie) {
        // Edit movie logic here
      } else {
        const newMovie = {
          title: this.title,
          director: this.director,
          year: this.year,
        };
        await createMovie(newMovie);
      }
      this.fetchMovies();
      this.clearForm();
    },
    editMovie(movie) {
      this.selectedMovie = movie;
      this.title = movie.title;
      this.director = movie.director;
      this.year = movie.year;
    },
    async deleteMovie(id) {
      await deleteMovie(id);
      this.fetchMovies();
    },
    clearForm() {
      this.title = '';
      this.director = '';
      this.year = null;
      this.selectedMovie = null;
    },
  },
  async created() {
    this.fetchMovies();
  },
};
</script>

<style scoped>
/* Add your styles here */
</style>
