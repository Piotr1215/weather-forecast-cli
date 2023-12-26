# Weather CLI Application

This command-line application retrieves and displays weather information for a specified location. The data includes current conditions, a three-day forecast, and sunrise/sunset times. It fetches data from the weatherapi.com service using an API key.

## Installation

To use the `justfile` commands, you must have `just` installed on your system. You can install it from [here](https://github.com/casey/just#installation).

To install the Weather CLI application, you need to have Go installed on your system. Follow these steps:

Alternatively, you can run the application using Docker. Ensure you have Docker installed and follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/Piotr1215/weather-forecast-cli
   ```

1. Build the Docker image:

   ```
   just build_container
   ```

1. Run the Docker container (make sure to replace `your_api_key_here` and `desired_location_here` with your actual API key and desired location):

   ```
   WEATHER_API_KEY=your_api_key_here WEATHER_LOCATION=desired_location_here just run_container
   ```

1. Build the application:
   ```
   go build -o weather-cli
   ```

## Usage

Before using the Weather CLI, you must set two environment variables: `WEATHER_API_KEY` and `WEATHER_LOCATION`. The API key can be obtained by signing up at weatherapi.com.

Set the environment variables:

```
export WEATHER_API_KEY=your_api_key_here
export WEATHER_LOCATION=desired_location_here
```

Run the application:

```
./weather-cli
```

The application will display the current weather conditions, a three-day forecast at 11 AM, and the sunrise/sunset times along with the duration of daylight for the specified location.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
