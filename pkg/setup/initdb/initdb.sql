\c db_nra_wsl;

CREATE TABLE release_manager_build_row (
    ProjectKey          VARCHAR PRIMARY KEY,
    BuildName           VARCHAR,
    Environment         VARCHAR,
    StartTimestamp      TIMESTAMP,
    CompletionTimestamp TIMESTAMP,
    Duration            float8,
    Phase               VARCHAR
);

CREATE TABLE result (
    ProjectKey         VARCHAR PRIMARY KEY,
    RepositoryName     VARCHAR,
    Technology         VARCHAR,
    ComponentType      VARCHAR,
    ODSVersion         VARCHAR,
    LastWebhookSuccess TIMESTAMP,
    ClusterRegion      VARCHAR
);

CREATE OR REPLACE VIEW result_set AS
    SELECT res.projectkey, res.repositoryname, res.technology, res.componenttype, res.odsversion, res.lastwebhooksuccess, res.clusterregion, rel.buildname, rel.environment, rel.starttimestamp, rel.completiontimestamp, rel.duration, rel.phase FROM result res 
        FULL OUTER JOIN release_manager_build_row rel
            ON res.ProjectKey = rel.ProjectKey; 


INSERT INTO release_manager_build_row (projectkey, buildname, environment, starttimestamp, completiontimestamp, duration, phase) VALUES ('project1', 'pro1', 'dev', '2004-10-19 10:23:54', '2005-10-19 11:11:11', '2222222', 'current');
INSERT INTO result (projectkey, repositoryname, technology, componenttype, odsversion, lastwebhooksuccess, clusterregion) VALUES ('project1', 'Relman', 'release', 'relman', '3.x', '2004-10-19 10:23:54', 'EU' );

INSERT INTO release_manager_build_row (projectkey, buildname, environment, starttimestamp, completiontimestamp, duration, phase) VALUES ('project2', 'pro2', 'dev', '2004-10-29 20:23:54', '2005-10-29 22:22:22', '12134234', 'current');
INSERT INTO result (projectkey, repositoryname, technology, componenttype, odsversion, lastwebhooksuccess, clusterregion) VALUES ('project2', 'Relman', 'release', 'relman', '3.x', '2004-10-29 20:23:54', 'US' );


