#include <iostream>
#include <curl/curl.h>
#include <string>
#include <chrono>
#include <random>
#include <ctime>
#include "nlohmann/json.hpp"
#include <conio.h>

using namespace std;
using json = nlohmann::json;

size_t write_callback(void* contents, size_t size, size_t nmemb, std::string* response) {
    size_t realsize = size * nmemb;
    response->append((char*)contents, realsize);
    return realsize;
}

std::string getCurrentTime() {
    auto now = std::chrono::system_clock::now();
    auto nowTime = std::chrono::system_clock::to_time_t(now);
    std::tm gmTimeInfo;
    gmtime_s(&gmTimeInfo, &nowTime);

    char buffer[80];
    std::strftime(buffer, 80, "%Y-%m-%d %H:%M:%S", &gmTimeInfo);
    return std::string(buffer);
}

void sendPostRequest(const std::string& url, const std::string& jsonData, const std::string& authToken) {
    CURL* curl;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_ALL);
    curl = curl_easy_init();
    if (curl) {
        struct curl_slist* headers = NULL;
        headers = curl_slist_append(headers, ("Authorization: Bearer " + authToken).c_str());
        headers = curl_slist_append(headers, "Content-Type: application/json");

        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
        curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, jsonData.c_str());

        std::string response_string;
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_string);

        res = curl_easy_perform(curl);
        if (res != CURLE_OK)
            std::cerr << "curl_easy_perform() failed: " << curl_easy_strerror(res) << std::endl;

        curl_easy_cleanup(curl);
        curl_slist_free_all(headers);
    }
    curl_global_cleanup();
}

std::string authenticateUser(const std::string& url, const std::string& username, const std::string& password) {
    CURL* curl;
    CURLcode res;
    std::string token;

    curl_global_init(CURL_GLOBAL_ALL);
    curl = curl_easy_init();
    if (curl) {
        json loginData = {
            {"username", username},
            {"password", password}
        };
        std::string jsonData = loginData.dump();

        curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, jsonData.c_str());

        std::string response_string;
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_string);

        res = curl_easy_perform(curl);
        if (res == CURLE_OK) {
            json responseJson = json::parse(response_string);
            if (responseJson.find("token") != responseJson.end()) {
                token = responseJson["token"].get<std::string>();
            }
            else {
                std::cerr << "Error: Token not found in response." << std::endl;
            }
        }
        else {
            std::cerr << "curl_easy_perform() failed: " << curl_easy_strerror(res) << std::endl;
        }

        curl_easy_cleanup(curl);
    }
    curl_global_cleanup();

    return token;
}

void simulateDataAndSend(const std::string& authToken) {
    std::string currentTime = getCurrentTime();
    std::random_device rd;
    std::mt19937 gen(rd());

    std::uniform_real_distribution<double> alcoholLevelDist(0.0, 0.5);
    double alcoholLevel = alcoholLevelDist(gen);

    std::string description = "Alcotest";

    json testResultData = {
        {"UserID", 1},
        {"TestTime", currentTime},
        {"AlcoholLevel", alcoholLevel},
        {"Description", description}
    };

    std::string jsonData = testResultData.dump();

    std::cout << "\n";
    std::cout << "Test Result Data:\n" << jsonData << std::endl;
    std::cout << "\n";

    std::string endpoint = "http://localhost:8000/api/user/1/testresults";

    sendPostRequest(endpoint, jsonData, authToken);
}

int main() {
    std::string username;
    std::string password;
    cout << "Enter username: ";
    cin >> username;
    cout << "Enter password: ";
    cin >> password;
    std::string authUrl = "http://localhost:8000/auth/login";

    std::string authToken = authenticateUser(authUrl, username, password);

    if (!authToken.empty()) {
        std::cout << "Authentication successful. Token obtained." << std::endl;

        bool stopSending = false;

        while (!stopSending) {
            simulateDataAndSend(authToken);
            Sleep(5000); 

            if (_kbhit()) {
                char key = _getch();
                if (key == 27) { 
                    stopSending = true;
                    std::cout << "\nSending stopped by user.\n";
                }
            }
        }
    }
    else {
        std::cerr << "Authentication failed. Unable to obtain token." << std::endl;
    }

    return 0;
}
