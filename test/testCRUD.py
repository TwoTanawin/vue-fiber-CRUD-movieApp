import requests
import json

BASE_URL = "http://localhost:3000/movies"

# Test the API for CRUD operations

# 1. Create a new movie
def create_movie():
    movie_data = {
        "title": "Inception",
        "director": "Christopher Nolan",
        "year": 2010
    }
    response = requests.post(BASE_URL, json=movie_data)
    if response.status_code == 200:
        print("Movie created successfully:", response.json())
        return response.json()["id"]  # Return the ID of the created movie for future tests
    else:
        print("Failed to create movie:", response.status_code, response.text)

# 2. Get all movies
def get_all_movies():
    response = requests.get(BASE_URL)
    if response.status_code == 200:
        print("Movies fetched successfully:", response.json())
    else:
        print("Failed to fetch movies:", response.status_code, response.text)

# 3. Get a single movie by ID
def get_movie(movie_id):
    response = requests.get(f"{BASE_URL}/{movie_id}")
    if response.status_code == 200:
        print(f"Movie with ID {movie_id} fetched successfully:", response.json())
    elif response.status_code == 404:
        print(f"Movie with ID {movie_id} not found")
    else:
        print("Failed to fetch movie:", response.status_code, response.text)

# 4. Update a movie
def update_movie(movie_id):
    updated_data = {
        "title": "The Dark Knight",
        "director": "Christopher Nolan",
        "year": 2008
    }
    response = requests.put(f"{BASE_URL}/{movie_id}", json=updated_data)
    if response.status_code == 200:
        print(f"Movie with ID {movie_id} updated successfully:", response.json())
    else:
        print("Failed to update movie:", response.status_code, response.text)

# 5. Delete a movie
def delete_movie(movie_id):
    response = requests.delete(f"{BASE_URL}/{movie_id}")
    if response.status_code == 200:
        print(f"Movie with ID {movie_id} deleted successfully")
    elif response.status_code == 404:
        print(f"Movie with ID {movie_id} not found")
    else:
        print("Failed to delete movie:", response.status_code, response.text)

if __name__ == "__main__":
    # Test creating a movie
    movie_id = create_movie()

    if movie_id:
        # Test fetching all movies
        get_all_movies()

        # Test fetching a single movie by ID
        get_movie(movie_id)

        # Test updating a movie
        update_movie(movie_id)

        # Test deleting a movie
        delete_movie(movie_id)

        # Fetch all movies again to ensure it's deleted
        get_all_movies()
