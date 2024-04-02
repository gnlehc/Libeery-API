SELECT *
FROM ms_lokers
WHERE Availability = 'Active' AND locker_id NOT IN (
    SELECT DISTINCT loker_id
    FROM tr_bookings
    WHERE session_id = 3
);