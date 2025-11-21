package co3017.microservices.content_service.config;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

/**
 * Main Application Class
 * Config cho Spring Boot application với Clean Architecture
 */
@SpringBootApplication(scanBasePackages = "co3017.microservices.content_service")
@EnableJpaRepositories(basePackages = "co3017.microservices.content_service.repository.postgresql")
@EntityScan(basePackages = "co3017.microservices.content_service.repository.postgresql.entity")
public class ContentServiceApplication {

	public static void main(String[] args) {
		SpringApplication.run(ContentServiceApplication.class, args);
	}

}

