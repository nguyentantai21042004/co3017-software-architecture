package co3017.microservices.content_service;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Main Application Class
 * Config cho Spring Boot application với Clean Architecture
 */
@SpringBootApplication(scanBasePackages = "co3017.microservices.content_service")
public class ContentServiceApplication {

	public static void main(String[] args) {
		SpringApplication.run(ContentServiceApplication.class, args);
	}

}
