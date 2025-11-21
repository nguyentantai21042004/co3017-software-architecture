package co3017.microservices.content_service.adapter.http;

import co3017.microservices.content_service.models.ContentVersion;
import co3017.microservices.content_service.usecase.ContentVersionUseCase;
import co3017.microservices.content_service.usecase.types.CreateContentVersionCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentVersionCommand;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.web.servlet.MockMvc;

import java.util.Optional;
import java.util.UUID;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

import co3017.microservices.content_service.adapter.http.response.ContentVersionResponseBuilder;

@WebMvcTest(controllers = ContentVersionController.class)
@ContextConfiguration(classes = { ContentVersionController.class, ContentVersionResponseBuilder.class })
@ActiveProfiles("test")
class ContentVersionControllerTest {

        @Autowired
        private MockMvc mockMvc;

        @MockBean
        private ContentVersionUseCase contentVersionUseCase;

        @Autowired
        private ObjectMapper objectMapper;

        @Test
        void create_ValidRequest_ReturnsCreated() throws Exception {
                UUID unitId = UUID.randomUUID();
                ObjectNode contentData = objectMapper.createObjectNode();
                contentData.put("text", "Sample content");

                CreateContentVersionCommand command = new CreateContentVersionCommand(unitId, "1.0.0", contentData,
                                true);
                ContentVersion createdVersion = new ContentVersion(1L, unitId, "1.0.0", contentData, true, null, null);

                when(contentVersionUseCase.create(any(CreateContentVersionCommand.class))).thenReturn(createdVersion);

                mockMvc.perform(post("/content-versions")
                                .contentType(MediaType.APPLICATION_JSON)
                                .content(objectMapper.writeValueAsString(command)))
                                .andExpect(status().isCreated())
                                .andExpect(jsonPath("$.version_id").value(1L))
                                .andExpect(jsonPath("$.version_number").value("1.0.0"));
        }

        @Test
        void getDetail_ExistingId_ReturnsVersion() throws Exception {
                Long id = 1L;
                ContentVersion version = new ContentVersion(id, UUID.randomUUID(), "1.0.0",
                                objectMapper.createObjectNode(),
                                true, null, null);

                when(contentVersionUseCase.detail(id)).thenReturn(Optional.of(version));

                mockMvc.perform(get("/content-versions/{id}", id))
                                .andExpect(status().isOk())
                                .andExpect(jsonPath("$.version_id").value(id));
        }

        @Test
        void getDetail_NonExistingId_ReturnsNotFound() throws Exception {
                Long id = 999L;
                when(contentVersionUseCase.detail(id)).thenReturn(Optional.empty());

                mockMvc.perform(get("/content-versions/{id}", id))
                                .andExpect(status().isNotFound());
        }

        @Test
        void update_ExistingId_ReturnsUpdatedVersion() throws Exception {
                Long id = 1L;
                UpdateContentVersionCommand command = new UpdateContentVersionCommand("1.1.0", null, true);
                ContentVersion updatedVersion = new ContentVersion(id, UUID.randomUUID(), "1.1.0",
                                objectMapper.createObjectNode(), true, null, null);

                when(contentVersionUseCase.update(eq(id), any(UpdateContentVersionCommand.class)))
                                .thenReturn(Optional.of(updatedVersion));

                mockMvc.perform(put("/content-versions/{id}", id)
                                .contentType(MediaType.APPLICATION_JSON)
                                .content(objectMapper.writeValueAsString(command)))
                                .andExpect(status().isOk())
                                .andExpect(jsonPath("$.version_number").value("1.1.0"));
        }
}
