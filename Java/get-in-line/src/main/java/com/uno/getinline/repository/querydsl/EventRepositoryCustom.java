package com.uno.getinline.repository.querydsl;

import com.uno.getinline.constant.EventStatus;
import com.uno.getinline.dto.EventViewResponse;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

import java.time.LocalDateTime;

public interface EventRepositoryCustom {
    Page<EventViewResponse> findEventViewPageBySearchParams(
            String placeName,
            String eventName,
            EventStatus eventStatus,
            LocalDateTime eventStartDatetime,
            LocalDateTime eventEndDatetime,
            // 페이징 데이타를 받을수 있다.
            Pageable pageable
    );
}
