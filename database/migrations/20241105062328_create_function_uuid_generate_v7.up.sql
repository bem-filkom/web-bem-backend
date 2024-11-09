CREATE OR REPLACE FUNCTION uuid_generate_v7()
    RETURNS uuid
AS
$$
    -- use random v4 uuid as starting point (which has the same variant we need)
    -- then overlay timestamp
    -- then set version 7 by flipping the 2 and 1 bit in the version 4 string
SELECT encode(
               set_bit(
                       set_bit(
                               overlay(uuid_send(gen_random_uuid())
                                       PLACING
                                       substring(int8send(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint)
                                                 FROM 3)
                                       FROM 1 FOR 6
                               ),
                               52, 1
                       ),
                       53, 1
               ),
               'hex')::uuid;
$$
    LANGUAGE SQL
    VOLATILE;
